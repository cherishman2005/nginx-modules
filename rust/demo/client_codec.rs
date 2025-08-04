use bcs::{Serialize, Deserialize};
use bytes::Bytes;
use anemo::Request;
use std::time::Duration;
use anyhow::{Context, Result};

// 定义BCS编解码器 trait
trait BcsCodec {
    /// 将类型序列化为BCS字节
    fn to_bcs(&self) -> Result<Bytes>
    where
        Self: Serialize,
    {
        bcs::to_bytes(self)
            .map(Bytes::from)
            .with_context(|| format!("BCS序列化{}失败", std::any::type_name::<Self>()))
    }

    /// 从BCS字节反序列化为类型
    fn from_bcs(bytes: &[u8]) -> Result<Self>
    where
        Self: Sized + Deserialize,
    {
        bcs::from_bytes(bytes)
            .with_context(|| format!("BCS反序列化{}失败", std::any::type_name::<Self>()))
    }
}

// 为请求和响应类型实现BCS编解码器
#[derive(Debug, Serialize, Deserialize)]
struct CommonRequest {
    action: String,
    data: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct CommonResponse {
    success: bool,
    result: String,
}

// 实现编解码器接口
impl BcsCodec for CommonRequest {}
impl BcsCodec for CommonResponse {}

// 发送请求（使用BCS编解码器）
pub async fn send_common_request(
    &self,
    peer_id: PeerId,
    route: &str,
    req: CommonRequest,
) -> Result<CommonResponse> {
    let network = self.network.read().await;
    let peer = network.peer(peer_id)
        .with_context(|| format!("节点 {} 未连接", peer_id))?;
    println!("客户端准备向服务路径 {} 发送请求", route);

    // 使用BCS编解码器序列化请求
    let req_bytes = req.to_bcs()?;
    println!("客户端请求体(BCS二进制长度): {}", req_bytes.len());
    
    let request = Request::new(req_bytes);
    
    // 发送请求并处理超时
    let response = tokio::time::timeout(
        Duration::from_secs(5),
        self.send_request(peer_id, route, request)
    )
    .await
    .map_err(|_| anyhow::anyhow!("请求超时"))?
    .map_err(|status| anyhow::anyhow!("请求失败: {:?}", status))?;

    let resp_body = response.body();
    println!("客户端收到响应体(BCS二进制长度): {}", resp_body.len());

    // 使用BCS编解码器反序列化响应
    let resp = CommonResponse::from_bcs(resp_body)?;
    Ok(resp)
}

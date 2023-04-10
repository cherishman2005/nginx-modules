# ingress

## 路径类型

Ingress 中的每个路径都需要有对应的路径类型（Path Type）。未明确设置 pathType 的路径无法通过合法性检查。当前支持的路径类型有三种：

* ImplementationSpecific：对于这种路径类型，匹配方法取决于 IngressClass。 具体实现可以将其作为单独的 pathType 处理或者与 Prefix 或 Exact 类型作相同处理。

* Exact：精确匹配 URL 路径，且区分大小写。

* Prefix：基于以 / 分隔的 URL 路径前缀匹配。匹配区分大小写，并且对路径中的元素逐个完成。 路径元素指的是由 / 分隔符分隔的路径中的标签列表。 如果每个 p 都是请求路径 p 的元素前缀，则请求与路径 p 匹配。


# 参考链接

- [https://kubernetes.io/zh-cn/docs/concepts/services-networking/ingress/](https://kubernetes.io/zh-cn/docs/concepts/services-networking/ingress/)

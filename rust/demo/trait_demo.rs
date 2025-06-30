trait ServerTrait {
    fn add(&self, a: i32, b: i32) -> i32;
}

#[derive(Default)]
struct Server{}

impl ServerTrait for Server {
	fn add(&self, a: i32, b: i32) -> i32 {
		return a+b;
	}
}

fn main() {
	let server = Server::default();
    println!("add={}", server.add(1,2));
}

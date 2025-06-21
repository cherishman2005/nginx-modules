fn main() {
	let opt = Some("hello");
	let val = match opt {
		Some(val) => val,
		None => "null",
	};
	println!("值是: {}", val);
}

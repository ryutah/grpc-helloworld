extern crate grpc;
extern crate protobuf;
extern crate tls_api;

mod helloworld;
mod helloworld_grpc;

use grpc::{RequestOptions, SingleResponse};
use helloworld::{HelloReply, HelloRequest};
use helloworld_grpc::{Greeter, GreeterServer};
use std::thread;

struct GreeterService;

impl Greeter for GreeterService {
    fn say_hello(&self, _opt: RequestOptions, req: HelloRequest) -> SingleResponse<HelloReply> {
        let name = if req.get_name().len() > 0 {
            req.get_name()
        } else {
            "no name"
        };
        println!("Receive name: {}", name);

        let mut reply = HelloReply::new();
        reply.set_message(format!("Hello, {}!!", name));

        SingleResponse::completed(reply)
    }
}

fn main() {
    let mut builder = grpc::ServerBuilder::new_plain();
    builder.http.set_port(8080);
    builder.add_service(GreeterServer::new_service_def(GreeterService));
    builder.build().expect("failed to build server");

    println!("Start server on port 8080\nPlease press Ctrl-C to stop server");

    loop {
        thread::park();
    }
}

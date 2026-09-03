use std::{
    io::{Read, Write},
    net::{TcpListener, TcpStream},
};

fn main() {
    println!("Hello world!");
    let lis = TcpListener::bind("localhost:8080").unwrap();
    for stream in lis.incoming() {
        println!("New conn...");
        println!("{:?}", stream);
        handle(stream.unwrap());
    }
}

fn handle(mut stream: TcpStream) {
    read_request(&mut stream);
    let status_line = "HTTP/1.1 200 OK";
    let contents = "<html><body><h1>Iam RUST!</h1></body></html>";
    let length = contents.len();
    let response = format!("{status_line}\r\nContent-Length: {length}\r\n\r\n{contents}");
    stream.write_all(response.as_bytes()).unwrap();
}

fn read_request(stream: &mut TcpStream) {
    let mut buf = [0; 1024];

    stream.read(&mut buf).unwrap();

    println!("Request: {}", String::from_utf8_lossy(&buf));
}

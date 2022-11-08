use serde::{Serialize, Deserialize};
// use warp::{self, http::Uri, path::FullPath, Filter, reply::Response, Reply };
use actix_web::{get, web, App, http::header, HttpServer, Responder, HttpResponse};
//use futures::future::{ok, Either, Ready};
// use rand::Rng;
//use std::collections::HashMap;
//use core::hash::Hash;
use std::str::FromStr;

#[derive(Deserialize)]
struct InputVideoUrl {
    video: String,
}

mod redirect_url;

#[get("/")]
async fn index(info: web::Query<InputVideoUrl>) -> HttpResponse {
    let cdn_host = "http://localhost:3200";
    let url = redirect_url::ab_redirect_by_rand(&info.video, &cdn_host); 
    HttpResponse::Found()
    .insert_header((header::LOCATION, url))
    .finish()
}

// #[get("/hello/{name}")]
// async fn greet(name: web::Path<String>) -> impl Responder {
//     format!("Hello {name}!")
// }

#[actix_web::main] // or #[tokio::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/hello", web::get().to(|| async { "Hello World!" }))
            //.service(greet)
            .service(index)
    })
    .bind(("0.0.0.0", 3200))?
    .run()
    .await
}

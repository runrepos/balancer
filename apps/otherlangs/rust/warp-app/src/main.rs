use serde::{Serialize, Deserialize};
use warp::{self, http::Uri, path::FullPath, Filter, reply::Response, Reply };
// use rand::Rng;
//use std::collections::HashMap;
//use core::hash::Hash;
use std::str::FromStr;

#[derive(Deserialize)]
struct InputVideoUrl {
    video: String,
}

mod redirect_url;


#[tokio::main]
async fn main() {

    // test
    let path = "http://s1.origin-cluster/video/1488/xcg2djHckad.m3u8";
    let cdn_host = "http://localhost:3200";

    let res = redirect_url::ab_redirect_by_rand(path, &cdn_host); //, &rng.gen::<f64>());
    println!("{}",res);

    let error404 = warp::path("error").map(|| "error 404");
    // let head_only = warp::head() //.map(warp::reply)
    // .map(|query: InputVideoUrl, p: FullPath| {
    //     resp(&query.video)
    // });
    let filter = warp::any()
        .and(warp::query::<InputVideoUrl>())
        .map(|i: InputVideoUrl| {
                let cdn_host = "http://localhost:3200";
                let url = redirect_url::ab_redirect_by_rand(&i.video, &cdn_host); //, rng.gen::<f64>()); //redirect_url::format_video_redirect_url(&cdn_host, &*i.video);
                warp::redirect(Uri::from_str(&url).unwrap()) 
            } );

        let routes = warp::get().and(
            error404
                .or(filter),
        ); //.or(head_only);

    warp::serve(routes)
        .run(([0, 0, 0, 0], 3200))
        .await;
}


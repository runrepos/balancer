// use serde::{Serialize, Deserialize};
#[macro_use] extern crate rocket;
use rocket::response::Redirect;
//use std::str::FromStr;

mod redirect_url;


 #[get("/?<video>")] //(&<color>")]
 fn hello(video: &str) -> Redirect { //, person: Person<'_>, other: Option<usize>) {
     let cdn_host = "http://localhost:3200";
     let url = redirect_url::ab_redirect_by_rand(&video, &cdn_host); 
     Redirect::moved(format!("{}", url)) // video
 }
 

 #[launch]
 fn rocket() -> _ {
     rocket::build().mount("/", routes![hello])
 }
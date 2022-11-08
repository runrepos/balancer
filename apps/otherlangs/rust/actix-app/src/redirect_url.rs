//std::ops::FnMut
use rand::Rng;

// #[derive(Debug)]
pub struct ServerUrl  {
    pub server: String,
    pub url: String,
}




pub fn ab_redirect_by_rand(path: &str, cdn_host: &str) -> String { //, ra: &f64) -> String {
    let mut rng = rand::thread_rng();
    if rng.gen::<f64>() < 0.3 {
        format_video_redirect_url(cdn_host, format_path(path))
    }else{ path.to_string() }
}

pub fn format_video_redirect_url(cdn_host: &str, redirect_url: ServerUrl) -> String { 
    // "/video/1488/xcg2djHckad.m3u8"
    cdn_host.to_owned()+"/"+&redirect_url.server+"/video/" +&redirect_url.url
}

pub fn format_path(path: &str) -> ServerUrl { //&str {
    //    let v: Vec<&str> = "abc1,def,2ghi".split(",").collect();
    //println!("{:?} days", v);
    // let v: Vec<&str> = path.split("http://").collect()[1];
    //let v: Vec<String> = vec!(path.split("http://").collect());

    let v: Vec<&str> = path.split("http://").collect();
    // println!("{}", &v[1]);
    let mut v: Vec<&str> = v[1].split(".").collect();
    let server: &str = &v[0];  v.remove(0); 
    let url: String = v.join("."); 
    let mut x: Vec<&str> = url.split("/").collect();
    x.remove(0); let joined_cutted_url = x.join("/");
    // println!("{:?}", x);
    // println!("{}", joined_cutted_url);
    let result = ServerUrl {
        server: server.to_string(),
        url: joined_cutted_url
    };
    return result ;
    
    //println!("{}", v[0]);
    //return "1";
}
#[macro_use]
extern crate rocket;
use rocket::fairing::{Fairing, Info, Kind};
use rocket::http::{ContentType, Header, Method, RawStr, Status};
use rocket::response::status::NoContent;
use rocket::serde::json::Json;
use rocket::{Request, Response};
use serde::{Deserialize, Serialize};

use rocket_cors::{AllowedOrigins, CorsOptions, AllowedMethods};

use serde_json::{self, json};

#[derive(Serialize, Deserialize)]
struct GameboardState {
    circles: Vec<String>,
    crosses: Vec<String>,
}

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[post("/updatedBoard", format = "application/json", data = "<data>")]
fn updated_board(data: Json<GameboardState>) -> serde_json::Value {
    // data.
    // let mut v: GameboardState = serde_json::from_str(&data).expect("Cannot parse json");
    let mut v = data;
    v.crosses.push("1,1".to_string());
    let circles = serde_json::to_string(&v.circles).expect("Cannot serialize json");
    let crosses = serde_json::to_string(&v.crosses).expect("Cannot serialize json");
    json!({ "status": "ok", "circles": circles, "crosses": crosses})
}
#[options("/updatedBoard", format="any")]
fn preflight2() -> Status {
     Status::Ok
}
#[options("/", format="any")]
fn preflight() -> Status {
     Status::Ok
}

pub struct CORS;

#[rocket::async_trait]
impl Fairing for CORS {
    fn info(&self) -> Info {
        Info {
            name: "Attaching CORS headers to responses",
            kind: Kind::Response
        }
    }

    async fn on_response<'r>(&self, _request: &'r Request<'_>, response: &mut Response<'r>) {
        response.set_header(Header::new("Access-Control-Allow-Origin", "*"));
        response.set_header(Header::new("Access-Control-Allow-Methods", "POST, GET, PATCH, OPTIONS"));
        response.set_header(Header::new("Access-Control-Allow-Headers", "*"));
        response.set_header(Header::new("Access-Control-Allow-Credentials", "true"));
    }
}

#[launch]
fn rocket() -> _ {
    // let cors = CorsOptions::default()
    //     .allowed_origins(AllowedOrigins::all())
    //     .allow_credentials(true);

    rocket::build()
        // .attach(cors.to_cors().unwrap().expect())
        .attach(CORS)
        .mount("/", routes![index])
}

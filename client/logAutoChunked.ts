import { API_URL } from "./consts";
import { logReadableStream } from "./logReadableStream";

export async function logAutoChunked() {
  const res = await fetch(API_URL + "/auto-chunk");
  console.log("Received headers");
  console.log("------------------- Start of headers -------------------");
  console.log(res.headers);
  console.log("------------------- End of headers -------------------\n");

  if (res.body) {
    logReadableStream(res.body);
  }
}

logAutoChunked();

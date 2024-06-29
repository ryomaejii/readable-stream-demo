const API_URL = "http://localhost:8080";

export async function logHeavyJsonStream(
  heavyJson: ReadableStream<Uint8Array>
) {
  const reader = heavyJson.getReader();
  let result = await reader.read();
  while (!result.done) {
    console.log("------------------- Start of chunk -------------------");
    console.log(`Chunk: ${new TextDecoder().decode(result.value)}`);
    console.log(
      `Character length: ${new TextDecoder().decode(result.value).length}`
    );
    console.log(`Byte length: ${result.value.byteLength}`);
    console.log("------------------- End of chunk -------------------\n");
    result = await reader.read();
  }
}

export async function fetchHeavyJson() {
  const res = await fetch(API_URL);
  console.log("Received headers");
  console.log("------------------- Start of headers -------------------");
  console.log(res.headers);
  console.log("------------------- End of headers -------------------\n");

  if (res.body) {
    logHeavyJsonStream(res.body);
  }
}

fetchHeavyJson();

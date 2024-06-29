export async function logReadableStream(stream: ReadableStream<Uint8Array>) {
  const reader = stream.getReader();
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

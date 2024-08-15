import { writable, type Writable } from "svelte/store";
import BACKEND_BASE_URL from "./config";

export const message = writable<MsgEvent>({ msg: "", eventtype: "" })
export const messages = writable<MsgEvent[]>([{ msg: "", eventtype: "" }])

type MsgTypes = (MsgEvent | MsgEvent[])

export type MsgEvent = {
  msg: string,
  eventtype: string,
}

export class Streamer {
  private url: string
  private controller: AbortController
  private streaming: Writable<MsgTypes>
  private arr: MsgEvent[] | undefined

  constructor(url: string, writableStore: Writable<MsgTypes>, array: MsgEvent[]) {
    this.url = url
    this.controller = new AbortController;
    this.streaming = writableStore
    this.arr = array
  }


  private readStream = (body: ReadableStream<Uint8Array>) => {
    const reader = body.getReader();
    const decoder = new TextDecoder();

    const read = async () => {
      const { done, value } = await reader.read();

      if (done) {
        return
      }

      const chunk = decoder.decode(value, { stream: true });
      let res: MsgEvent = JSON.parse(chunk)
      if (this.arr) {
        console.log(this.arr)
        return read();
      }
      this.streaming.set(res)
      return read();
    }

    return read();
  }

  connect = async () => {
    console.log(this.url)
    const response = await fetch(BACKEND_BASE_URL + this.url, {
      signal: this.controller.signal
    });

    try {
      if (response && response.ok && response.body) {
        await this.readStream(response.body);
      }
    }

    catch (err) {
      if (err instanceof TypeError) {
        console.log("typeerror: ", err.message)
      }
      else if (err instanceof DOMException) {
        console.log(err.message)
      } else {
        console.log("Unhandled error: ", err)
      }
    }

    finally {
      this.controller = new AbortController;
    }
  }

  disconnect = () => {
    this.controller.abort("disconnect")
  }
}

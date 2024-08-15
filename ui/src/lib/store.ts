import { writable } from 'svelte/store';
import type { Hello } from '$lib/types'
import { type MsgEvent } from './stream';

export const apiData = writable<Hello>({ "hello": "" });
export const timeMessages = writable<MsgEvent>({ msg: "", eventtype: "" })
export const memoryMessages = writable<MsgEvent>({ msg: "", eventtype: "" })
export const syslogMessages = writable<MsgEvent[]>([{ msg: "", eventtype: "" }])
export const routeLoadingState = writable<boolean>(false);

import { writable } from 'svelte/store';
import type { Hello } from '$lib/api'

export const apiData = writable<Hello>({ "hello": "" });

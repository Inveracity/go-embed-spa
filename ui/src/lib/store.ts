import { writable } from 'svelte/store';
import type { Hello } from '$lib/types'

export const apiData = writable<Hello>({ "hello": "" });

export const routeLoadingState = writable<boolean>(false);

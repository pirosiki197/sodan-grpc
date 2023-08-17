import { PromiseClient, createPromiseClient } from "@bufbuild/connect";
import {
    createConnectTransport,
  } from "@bufbuild/connect-web";
import { ServiceType } from "@bufbuild/protobuf";

const url = import.meta.env.VITE_API_URL ? import.meta.env.VITE_API_URL : "https://back.sodan.trap.show";

const transport = createConnectTransport({
    baseUrl: url,
})

export function useClient<T extends ServiceType>(service: T): PromiseClient<T> {
    const client = createPromiseClient(service, transport);
    return client;
}
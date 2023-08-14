import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { APIService } from "./api/pb/api/v1/api_connect";

const url = import.meta.env.VITE_API_URL ? import.meta.env.VITE_API_URL : "https://back.sodan.trap.show";

const transport = createConnectTransport({
    baseUrl: url,
})

const client = createPromiseClient(APIService, transport);

export default client;
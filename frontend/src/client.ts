import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { APIService } from "./api/pb/api/v1/api_connect";

const url = "http://localhost:8080";

const transport = createConnectTransport({
    baseUrl: url,
})

const client = createPromiseClient(APIService, transport);

export default client;
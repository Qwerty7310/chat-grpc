import {ChatServiceClient} from "../generated/ChatServiceClientPb";

export const chatClient = new ChatServiceClient("http://localhost:8081", null, {
    withCredentials: false,
});
import {chatClient} from "./grpc.ts"; // ./grpc
import { MessageToServer, HistoryRequest} from "../generated/chat_pb";

export async function sendMessage(text: string): Promise<void> {
    const req = new MessageToServer();
    req.setText(text)

    await chatClient.sendMessage(req, {})
}

export async function getHistory(limit: number) {
    const req = new HistoryRequest();
    req.setLimit(limit);

    return await chatClient.getHistory(req, {})
}
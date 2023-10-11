import {v4 as uuidv4} from 'uuid';
import {config} from "./config.js";


export function createMapperHeader(mode) {
    return {
        "version": "0.1.0",
        "message_id": uuidv4(),
        "message_ts": new Date().toISOString(),
        "action": mode,
        "sender_id": "mapper.sunbirdrc.dev",
        "sender_uri": "",
        "receiver_id": "callback.sunbirdrc.dev",
        "total_count": 0,
        "is_encrypted": false
    }
}

const TRANSACTION_ID = "transaction_id";

export function createTransactionObject() {
    let existingTransactionID = localStorage.getItem(TRANSACTION_ID);
    if (!existingTransactionID) {
        existingTransactionID = uuidv4();
        localStorage.setItem(TRANSACTION_ID, existingTransactionID)
    }
    return {
        transaction_id: existingTransactionID,
        reference_id: uuidv4()
    }
}

export function getTransactionId() {
    return localStorage.getItem(TRANSACTION_ID);
}

import api from "../api.ts";


export const homeCardsService = () => {
    return api.get("/home/cards");
}
import api from "@/services/api.ts";


export const homeCardsService = () => {
    return api.get("/home/cards");
}
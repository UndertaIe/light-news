import axios from "axios";
import Constants from "./consts";

const searchUrl = `${Constants.Domain}/api/v1/search`;

const SearchService = {
  search: (content: string, pagination: any) => {
    const params = new URLSearchParams(pagination);
    return axios.get(`${searchUrl}/${content}?${params}`);
  }
};

export default SearchService;
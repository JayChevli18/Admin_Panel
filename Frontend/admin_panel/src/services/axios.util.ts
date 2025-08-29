import axios, { AxiosRequestConfig, AxiosResponse } from "axios";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export const axiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET, POST, PUT, PATCH, DELETE, OPTIONS",
  },
});

export const AsyncFetch = <T>(
  request: AxiosRequestConfig & {
    next?: { revalidate: number };
    fullResponse?: boolean;
  }
): Promise<T> => {
  return new Promise(async (resolve, reject) => {
    axiosInstance(request)
      .then(({ data, error }: any) => {
        if (request.fullResponse) {
          resolve(data);
        } else {
          if (data?.success) {
            resolve(data.data);
          } else {
            reject(data?.error || "Something went wrong");
            throw new Error(error);
          }
        }
      })
      .catch((error) => {
        reject(error);
      });
  });
};

axiosInstance.interceptors.response.use((response)=>{
    return response;
},(error)=>{
    return Promise.reject(error);
})
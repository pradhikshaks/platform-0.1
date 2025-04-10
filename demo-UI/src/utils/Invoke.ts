import axios from "axios";


const getData = (endpoint: string) => {
  const baseurl = import.meta.env.VITE_API_URL;
  const url = `${baseurl}/data/${endpoint}`;  
  // const data = fetch(url).then((res)=> res.json()).catch((err)=> console.log(err))
  return axios.get(url)
    .then((response) => {
      return response.data.Data;
    })
    .catch((error) => {
      console.error('Error fetching data:', error);
      return [];
    });
  // return data
};

export {getData}
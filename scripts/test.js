import http from 'k6/http';

export let options = {
    vus: 10,
    // duration: '1m',
    iterations: 200000 
}
export default function () {
    const url = 'http://host.docker.internal:9000/v1';
    const payload = JSON.stringify({
        name:"Kub",
        side:true,
        price:500,
        amount:51.55
    });
  
    const params = {
      headers: {
        'Content-Type': 'application/json',
      },
    };
  
    http.post(url, payload, params);
}

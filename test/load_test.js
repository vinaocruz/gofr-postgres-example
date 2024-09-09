import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 200 },
    { duration: '30s', target: 2000 },
    { duration: '2m', target: 0 },
  ],
};

export default function() {
  http.get('http://app:8000/books');
  sleep(1);
}

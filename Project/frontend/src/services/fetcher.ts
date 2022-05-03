import axios from 'axios';

const fetcher = axios.create({
    baseURL: "https://mr-maze-backend-game-dev.herokuapp.com",
    headers: {
      'Content-Type': 'application/json',
    },
});

export default fetcher
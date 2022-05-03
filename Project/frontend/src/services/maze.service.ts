import fetcher from "./fetcher";

const MazeService = {
    generate: () => fetcher.post('/generate', { height: 20, width: 20 }),
}

export default MazeService;
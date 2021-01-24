import * as dotenv from 'dotenv';
import express from 'express';
import { createConnection } from 'typeorm';

dotenv.config();
const app = express();
app.use(express.json());

createConnection();

export default app;

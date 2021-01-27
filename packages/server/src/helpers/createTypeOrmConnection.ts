import { createConnection, getConnectionOptions } from 'typeorm';

const createTypeOrmConnection = async () => {
  const connectionOptions = await getConnectionOptions(process.env.NODE_ENV);
  return createConnection({ ...connectionOptions, name: 'default' });
};

export default createTypeOrmConnection;

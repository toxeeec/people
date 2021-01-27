import createTypeOrmConnection from './helpers/createTypeOrmConnection';
import app from './server';

const { PORT } = process.env;

createTypeOrmConnection().then(() => {
  app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
  });
});

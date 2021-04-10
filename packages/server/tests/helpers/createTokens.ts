import dotenv from 'dotenv';
import faker from 'faker';
import jwt from 'jsonwebtoken';
import { getConnection } from 'typeorm';
import { User } from '../../src/entity/User';
import createTokens from '../../src/helpers/createTokens';
import createTypeOrmConnection from '../../src/helpers/createTypeOrmConnection';

dotenv.config({ path: '.env.test' });

const sampleUser = {
  name: faker.name.firstName(),
  surname: faker.name.lastName(),
  email: faker.internet.email(),
  password: faker.internet.password(),
};

beforeAll(async () => {
  await createTypeOrmConnection();
});

describe('createTokens helper function', () => {
  it('should return accessToken and refreshToken', async (done) => {
    const user = await User.create(sampleUser).save();
    const { accessToken, refreshToken } = await createTokens(user);
    const decodedAccessToken = jwt.verify(
      accessToken,
      process.env.ACCESS_TOKEN_SECRET!
    );
    const decodedRefreshToken = jwt.verify(
      refreshToken,
      process.env.REFRESH_TOKEN_SECRET!
    );
    const { id, name, surname } = user;
    expect(decodedAccessToken).toContainValues([id, name, surname]);
    expect(decodedRefreshToken).toContainValue(id);
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});

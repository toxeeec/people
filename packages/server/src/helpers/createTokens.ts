import jwt from 'jsonwebtoken';
import { RefreshToken } from '../entity/RefreshToken';
import { User } from '../entity/User';

const createTokens = async (user: User) => {
  const accessToken = jwt.sign(
    { id: user.id, name: user.name, surname: user.surname },
    process.env.ACCESS_TOKEN_SECRET,
    { expiresIn: '15m' }
  );
  const refreshToken = jwt.sign(
    { id: user.id },
    process.env.REFRESH_TOKEN_SECRET
  );

  await RefreshToken.create({
    user,
    token: refreshToken,
  }).save();
  return { accessToken, refreshToken };
};

export default createTokens;

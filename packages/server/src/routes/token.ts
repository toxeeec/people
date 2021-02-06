import { Router } from 'express';
import jwt from 'jsonwebtoken';
import { RefreshToken } from '../entity/RefreshToken';
import ApiError from '../helpers/ApiError';

const router = Router();

router.post('/token', async (req, res, next) => {
  const authHeader = req.headers.authorization;
  const token = authHeader && authHeader.split(' ')[1];
  if (!token) return next(ApiError.unauthorized());
  let userId: string;
  jwt.verify(
    token,
    process.env.REFRESH_TOKEN_SECRET,
    (err, user: { id: string }) => {
      if (err) return next(ApiError.forbidden());
      userId = user.id;
    }
  );
  try {
    const validToken = await RefreshToken.findOne({ token });
    if (!validToken) return next(ApiError.forbidden());
    const newToken = jwt.sign({ id: userId }, process.env.ACCESS_TOKEN_SECRET, {
      expiresIn: '15m',
    });
    res.json({ accessToken: newToken });
  } catch (err) {
    next(ApiError.internal());
  }
});

export default router;

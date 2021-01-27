import * as yup from 'yup';

export const authenticateUserSchema = yup.object().shape({
  email: yup.string().required().email(),
  password: yup.string().required(),
});

export const registerUserSchema = yup.object().shape({
  name: yup.string().required().trim(),
  surname: yup.string().required().trim(),
  email: yup.string().required().email(),
  password: yup.string().required().min(8),
});

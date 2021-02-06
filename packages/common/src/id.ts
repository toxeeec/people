import * as yup from 'yup';

export const idSchema = yup.object().shape({
  id: yup.string().required().uuid(),
});

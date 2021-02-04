import * as yup from 'yup';

export const validateIdSchema = yup.object().shape({
  id: yup.string().required().uuid(),
});

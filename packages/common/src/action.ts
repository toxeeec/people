import * as yup from 'yup';

export const frienRequestActionSchema = yup.object().shape({
  action: yup.string().required().oneOf(['accept', 'decline']),
});

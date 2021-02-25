import * as yup from 'yup';

export const friendRequestActionSchema = yup.object().shape({
  action: yup.string().required().oneOf(['accept', 'decline']),
});

import * as yup from 'yup';

export const paginateSchema = yup.object().shape({
  page: yup.number().positive().notRequired().default(1),
  limit: yup.number().positive().notRequired(),
});

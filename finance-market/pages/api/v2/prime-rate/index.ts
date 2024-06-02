import { NextApiRequest, NextApiResponse } from 'next';
import { handleInflationOrPrimeRate } from '~/server/api/handleInflationOrPrimeRate';

export default async (req: NextApiRequest, res: NextApiResponse) => {
  return handleInflationOrPrimeRate('prime-rate', req, res);
};

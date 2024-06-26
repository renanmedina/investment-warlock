export const numberToSI = (number: any) => {
  var SI_SYMBOL = ['', 'k', 'M', 'B', 'T'];

  var tier = (Math.log10(Math.abs(number)) / 3) | 0;

  if (tier == 0) return number;

  var suffix = SI_SYMBOL[tier];
  var scale = Math.pow(10, tier * 3);

  var scaled = number / scale;

  return scaled.toFixed(1) + suffix;
};

export const numberToMoney = (number: any) => {
  const parsedNumber = parseFloat(number);

  if (isNaN(parsedNumber)) return '--';

  let money = new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
  });

  return money.format(parsedNumber);
};

export const numberToSIMoney = (number: any) => {
  const parsedNumber = parseFloat(number);

  if (isNaN(parsedNumber)) return '--';

  let money = new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
    notation: 'compact',
  });

  return money.format(parsedNumber);
};

export const numberToPercent = (number: any) => {
  let percent = new Intl.NumberFormat('pt-BR', {
    style: 'percent',
    maximumSignificantDigits: 2,
  });

  return percent.format(number / 100);
};

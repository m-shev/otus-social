import * as React from 'react';
import {useDots} from './useDots';

export const SmallDots: React.FC = () => {
    const dots = useDots();
    return <div>{dots}</div>;
};

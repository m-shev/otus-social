import * as React from 'react';
import {useDots} from './useDots';

export type SmallDotsProps = {};

export const SmallDots: React.FC<SmallDotsProps> = () => {
    const dots = useDots();
    return <div>{dots}</div>;
};

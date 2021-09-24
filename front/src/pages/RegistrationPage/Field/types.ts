import {FieldInputProps} from 'formik';

export type FieldSelectOption = {
    value: string;
    title: string;
};

export type BaseFieldProps = FieldInputProps<
    string | ReadonlyArray<string> | number | undefined
> & {
    id: string;
    title: string;
    type: string;
    required: boolean;
};

export type SelectField = BaseFieldProps & {
    options: FieldSelectOption[];
    type: 'select';
};

export type NumberField = BaseFieldProps & {
    min?: number;
    max?: number;
    type: 'number';
};

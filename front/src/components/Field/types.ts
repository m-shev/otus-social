import {FieldInputProps} from 'formik';

export enum FieldType {
    Select = 'select',
    Input = 'input',
    Checkbox = 'checkbox',
}

export type FieldOption = {
    value: string;
    title: string;
};

export type CheckboxFieldOptions = FieldOption & {
    id: string;
};

export type BaseFieldProps = FieldInputProps<
    string | ReadonlyArray<string> | number | undefined
> & {
    error?: string | string[];
    id: string;
    title: string;
    type: FieldType | string;
    required: boolean;
    setFieldValue: (value: any) => void;
};

export type SelectFieldProps = BaseFieldProps & {
    options: FieldOption[];
    type: 'select';
};

export type CheckboxFieldProps = BaseFieldProps & {
    options: FieldOption[];
    type: 'checkbox';
};

export type NumberFieldProps = BaseFieldProps & {
    min?: number;
    max?: number;
    type: 'number';
};

export type FieldProps = BaseFieldProps | SelectFieldProps | NumberFieldProps | CheckboxFieldProps;

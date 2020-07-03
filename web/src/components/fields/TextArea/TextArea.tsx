/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import React from 'react';
import { Box, FormError, TextArea, TextAreaProps } from 'pouncejs';
import { FieldConfig, useField } from 'formik';

const FormikTextArea: React.FC<TextAreaProps & Required<Pick<FieldConfig, 'name'>>> = props => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [field, meta] = useField(props.name);

  const isInvalid = meta.touched && !!meta.error;
  const errorElementId = isInvalid ? `${props.name}-error` : undefined;
  return (
    <Box>
      <TextArea
        invalid={isInvalid}
        aria-describedby={isInvalid ? errorElementId : undefined}
        {...props}
      />
      {isInvalid && (
        <FormError mt={2} id={errorElementId}>
          {meta.error}
        </FormError>
      )}
    </Box>
  );
};

export default FormikTextArea;

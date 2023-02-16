// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Provider for FortiOS

package fortiadc

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func set_multipart_parameter(writer *multipart.Writer, key, data string, is_file bool) error {
	fw, err := writer.CreateFormField(key)
	if err != nil {
		return err
	}

	if is_file {
		if _, err := os.Stat(data); err != nil {
			return err
		}
		file, err := os.Open(data)
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			return err
		}
	} else {
		_, err = io.Copy(fw, strings.NewReader(data))
		if err != nil {
			return err
		}
	}
	return nil
}

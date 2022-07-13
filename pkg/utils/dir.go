/*
 * Copyright 2022 Xiongfa Li.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Rmdir(target string) error {
	info, err := os.Stat(target)
	if err != nil {
		return nil
	} else if info.IsDir() {
		return os.RemoveAll(target)
	} else {
		return fmt.Errorf("Generate failed: Target %s is not a dir. ", target)
	}
	return nil
}

func Mkdir(target string) error {
	info, err := os.Stat(target)
	if err != nil {
		err = os.MkdirAll(target, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Generate failed: create target dir %s error: %v. ", target, err)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("Generate failed: Target %s is not a dir. ", target)
	}
	return nil
}

func CopyDir(dst, src string) error {
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("file: %s not exists. ", src)
	}
	if info.IsDir() {
		return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			r, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}
			dstPath := filepath.Join(dst, r)
			if info.IsDir() {
				return Mkdir(dstPath)
			} else {
				_, err = CopyFile(dstPath, path)
				return err
			}
		})
	} else {
		_, err := CopyFile(dst, src)
		if err != nil {
			return fmt.Errorf("Copy file %s to %s failed. ", src, dst)
		}
	}
	return nil
}

func CopyFile(dst, src string) (int64, error) {
	info, err := os.Stat(src)
	if err != nil {
		return 0, fmt.Errorf("File: %s not exists. ", src)
	} else if info.IsDir() {
		return 0, fmt.Errorf("File: %s is a dir. ", src)
	}
	_, err = os.Stat(dst)
	if err == nil {
		return 0, fmt.Errorf("File: %s is exists. ", dst)
	}
	dstf, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, info.Mode())
	if err != nil {
		return 0, fmt.Errorf("Create dst file: %s failed. %v ", dst, err)
	}
	defer dstf.Close()
	srcf, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("Open src file: %s failed. %v ", src, err)
	}
	defer srcf.Close()
	return io.Copy(dstf, srcf)
}

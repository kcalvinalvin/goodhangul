/*
goodhangul
Copyright (C) 2020 Calvin Kim

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
hwp50 package implements the hwp file version 5.0 based on the specification
from Hancom.

hwp50 패키지는 한컴에서 제공한 hwp 파일 버전 5.0을 구현한 것입니다.

Hwp file struct looks like such

Hwp 파일 구조는 이와 같습니다

	FileHeader
	DocInfo
	BodyText:
		Section0
		Section1
		...
	SummaryInfo
	BinaryData:
		BinaryData0
		BinaryData1
		...
	PreviewText
	PreviewImage
	DocOptions:
		_LinkDoc
		DrmLicense
		...
	Scripts:
		DefaultJScript
		JScriptVersion
		...
	XMLTemplate:
		Schema
		Instance
		...
	DocHistory
		VersionLog0
		VersionLog1
		...

Each must be read
*/
package hwp50

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
from Hancom. All the encoding is done in UTF-16LE.

Hwp file struct looks like such

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

///////////////////////////////////////////////////////////////////////////////
// COPYRIGHT (c) 2014 Schweitzer Engineering Laboratories, Inc.
//
// This file is provided under a BSD license. The text of the BSD license 
// is provided below.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. The name of the author may not be used to endorse or promote products
// derived from this software without specific prior written permission.
//
// Alternatively, provided that this notice is retained in full, this software
// may be distributed under the terms of the GNU General Public License ("GPL")
// version 2, in which case the provisions of the GPL apply INSTEAD OF those
// given above.
//
// This program is free software; you can redistribute it and/or modify it
// under the terms and conditions of the GNU General Public License,
// version 2, as published by the Free Software Foundation.
//
// This program is distributed in the hope it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin St - Fifth Floor, Boston, MA 02110-1301 USA.
//
// The full GNU General Public License is included in this distribution in
// the file called "COPYING".
//
// Contact Information:
// SEL Opensource <opensource@selinc.com>
// Schweitzer Engineering Laboratories, Inc.
// 2350 NE Hopkins Court, Pullman, WA 99163
///////////////////////////////////////////////////////////////////////////////
/// @brief
/// Linux and Windows drivers utilize different standard types. This file
/// allows common files between them to compile clean, without having to make
/// updates to types.
///////////////////////////////////////////////////////////////////////////////
#ifndef SEL_STD_TYPES_H_INCLUDED
#define SEL_STD_TYPES_H_INCLUDED

#ifdef __cplusplus
extern "C"
{
#endif

#if defined(__KERNEL__) // Defined under Linux

#include <linux/types.h>

#define UINT64 u64
#define UINT32 u32
#define UINT16 u16
#define UINT8  u8

#define VOID void
#define BOOLEAN u8
#define TRUE 1
#define FALSE 0

#define ASSERT(x) BUG_ON(!(x))

#elif defined(_MSC_VER) // defined under Windows

#include <basetsd.h>

#define u64  UINT64
#define u32  UINT32
#define u16  UINT16
#define u8   UINT8

#endif

#ifdef __cplusplus
}
#endif

#endif // SEL_STD_TYPES_H_INCLUDED

/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
#pragma once

#include <czmq.h>

#ifdef __cplusplus
extern "C" {
#endif

#include "intertask_interface.h"
#include "mme_app_ue_context.h"
#include "nas_timer.h"

#define MME_APP_TIMER_INACTIVE_ID (-1)

int mme_app_start_timer_arg(
    size_t msec, timer_repeat_t repeat, zloop_timer_fn handler,
    timer_arg_t* arg);

// Most handlers only need mme_ue_s1ap_id, use this function for
// such handlers.
int mme_app_start_timer(
    size_t msec, timer_repeat_t repeat, zloop_timer_fn handler,
    mme_ue_s1ap_id_t ue_id);

void mme_app_stop_timer(int timer_id);

void mme_app_resume_timer(
    struct ue_mm_context_s* const ue_mm_context_pP, time_t start_time,
    nas_timer_t* timer, zloop_timer_fn timer_expiry_handler, char* timer_name);

bool mme_app_get_timer_arg(int timer_id, timer_arg_t* arg);
bool mme_app_get_timer_arg_ue_id(int timer_id, mme_ue_s1ap_id_t* ue_id);

#ifdef __cplusplus
}
#endif

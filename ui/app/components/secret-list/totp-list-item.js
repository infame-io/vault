/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';

export default class TotpListItemComponent extends Component {
  @tracked showConfirmModal = false;
}

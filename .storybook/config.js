import { configure, addDecorator } from '@storybook/react';
import { withInfo } from '@storybook/addon-info';
import 'loki/configure-react';

import '../src/index.scss';

function loadStories() {
  // TODO automatically scan stories dir see react-uswds repo for example
  require('../src/stories/index.stories.jsx');
  require('../src/stories/statusTimeLine.stories.jsx');
  require('../src/stories/tabNav.stories.jsx');
}

addDecorator(withInfo);
configure(loadStories, module);

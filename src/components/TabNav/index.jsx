import React from 'react';
import PropTypes from 'prop-types';
import { Tag } from '@trussworks/react-uswds';
import { Tab, Tabs, TabList } from 'react-tabs';
import classNames from 'classnames/bind';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

const TabNav = ({ options, children }) => (
  <Tabs className={cx('tab-nav')}>
    <div className={cx('tab-list-container')}>
      <TabList className={cx('tab-list')}>
        {options.map(({ title, url, notice }, index) => (
          <Tab key={index.toString()} selectedClassName={cx('tab-active')}>
            <a href={url} className={cx({ 'tab-with-notice': notice })}>
              <span className={cx('tab-title')}>{title}</span>
            </a>
            {notice && <Tag>{notice}</Tag>}
          </Tab>
        ))}
      </TabList>
    </div>
    {children}
  </Tabs>
);

TabNav.propTypes = {
  options: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string,
      url: PropTypes.string,
      notice: PropTypes.string,
    }),
  ).isRequired,
  children: PropTypes.arrayOf(PropTypes.node).isRequired,
};

export default TabNav;

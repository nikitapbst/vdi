package main

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"log"
	"net/http"

	"github.com/kolo/xmlrpc"
	"github.com/megamsys/libgo/cmd"

	"github.com/gotk3/gotk3/gtk"
	"github.com/megamsys/opennebula-go/api"
	"github.com/zalando/go-keyring"
)

const morda string = `
<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.2 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkDialog" id="AdminCheckDialog">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Пароль администратора</property>
    <property name="type_hint">dialog</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="acdNext">
                <property name="label" translatable="yes">Далее</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="acdCancel">
                <property name="label" translatable="yes">Отмена</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkGrid">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="row_spacing">3</property>
            <property name="column_spacing">3</property>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">Пароль:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="rememberedPas">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkDialog" id="AuthorizationDialog">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Авторизация</property>
    <property name="type_hint">dialog</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="authDnext">
                <property name="label" translatable="yes">Далее</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkGrid">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="row_spacing">3</property>
            <property name="column_spacing">3</property>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Логин:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Пароль:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="LoginEntry">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="PasswordEntry">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="visibility">False</property>
                <property name="invisible_char">*</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkDialog" id="ErrorDialog">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Ошибка</property>
    <property name="type_hint">dialog</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">3</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="errOk">
                <property name="label" translatable="yes">Ок</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <signal name="clicked" handler="on_errOk_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="label" translatable="yes">Ошибка:</property>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="errlbl">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkDialog" id="StartDialog">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">Окно администратора</property>
    <property name="type_hint">dialog</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="startNext">
                <property name="label" translatable="yes">Далее</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <signal name="clicked" handler="on_startNext_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="startCancel">
                <property name="label" translatable="yes">Отмена</property>
                <property name="visible">True</property>
                <property name="sensitive">False</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkGrid">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="row_spacing">3</property>
            <property name="column_spacing">3</property>
            <child>
              <object class="GtkEntry" id="portEntry">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="brokerEntry">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <signal name="changed" handler="on_brokerEntry_changed" swapped="no"/>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Брокер соединения:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkCheckButton" id="saveCheckButton">
                <property name="label" translatable="yes">Запомнить</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">False</property>
                <property name="draw_indicator">True</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">4</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Порт:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="adminPass">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="visibility">False</property>
                <property name="invisible_char">*</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="repeatAdmPass">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="visibility">False</property>
                <property name="invisible_char">*</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Пароль администратора:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Повторить пароль:</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">3</property>
              </packing>
            </child>
            <child>
              <placeholder/>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkListStore" id="tempVmList">
    <columns>
      <!-- column-name name -->
      <column type="gchararray"/>
      <!-- column-name color -->
      <column type="gchararray"/>
    </columns>
  </object>
  <object class="GtkListStore" id="vmList">
    <columns>
      <!-- column-name name -->
      <column type="gchararray"/>
      <!-- column-name state -->
      <column type="gchararray"/>
      <!-- column-name protocol -->
      <column type="gchararray"/>
      <!-- column-name id -->
      <column type="gchararray"/>
    </columns>
  </object>
  <object class="GtkDialog" id="MachineControlDialog">
    <property name="can_focus">False</property>
    <property name="resizable">False</property>
    <property name="default_width">450</property>
    <property name="default_height">300</property>
    <property name="type_hint">dialog</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="turnOnOff">
                <property name="label" translatable="yes">Включить</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">start</property>
                <property name="hexpand">True</property>
                <signal name="clicked" handler="on_turnOnOff_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="reboot">
                <property name="label" translatable="yes">Перезагрузить</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">center</property>
                <signal name="clicked" handler="on_reboot_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="connectTo">
                <property name="label" translatable="yes">Соединится</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">end</property>
                <signal name="clicked" handler="on_connectTo_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkButtonBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="halign">end</property>
            <property name="spacing">2</property>
            <property name="layout_style">start</property>
            <child>
              <object class="GtkButton" id="reAuth">
                <property name="label" translatable="yes">Переавторизация</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">start</property>
                <property name="resize_mode">queue</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="turnOffComp">
                <property name="label" translatable="yes">Выключение компьютера</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">end</property>
                <signal name="clicked" handler="on_turnOffComp_clicked" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkScrolledWindow">
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="shadow_type">in</property>
            <child>
              <object class="GtkTreeView" id="vmListTree">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="model">vmList</property>
                <child internal-child="selection">
                  <object class="GtkTreeSelection" id="vmSelection"/>
                </child>
                <child>
                  <object class="GtkTreeViewColumn" id="name_of_vm">
                    <property name="resizable">True</property>
                    <property name="title" translatable="yes">Имя</property>
                    <property name="expand">True</property>
                    <child>
                      <object class="GtkCellRendererText" id="namevm"/>
                      <attributes>
                        <attribute name="background-gdk">0</attribute>
                        <attribute name="foreground-gdk">0</attribute>
                        <attribute name="text">0</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
                <child>
                  <object class="GtkTreeViewColumn" id="state_of_vm">
                    <property name="resizable">True</property>
                    <property name="title" translatable="yes">Состояние</property>
                    <property name="expand">True</property>
                    <child>
                      <object class="GtkCellRendererText" id="statevm"/>
                      <attributes>
                        <attribute name="text">1</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
                <child>
                  <object class="GtkTreeViewColumn" id="VmId">
                    <property name="visible">False</property>
                    <property name="title" translatable="yes">ID</property>
                    <child>
                      <object class="GtkCellRendererText" id="VmIdValue"/>
                      <attributes>
                        <attribute name="text">3</attribute>
                      </attributes>
                    </child>
                  </object>
                </child>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>


`

type VmState int
type LcmState int

var LcmStateString = map[LcmState]string{LCM_INIT: "lcm_init", PROLOG: "prolog", BOOT: "boot", RUNNING: "running",
	MIGRATE: "migrate", SAVE_STOP: "save_stop", SAVE_SUSPEND: "save_suspend", SAVE_MIGRATE: "save_migrate",
	PROLOG_MIGRATE: "prolog_migrate", PROLOG_RESUME: "prolog_resume", EPILOG_STOP: "epilog_stop", EPILOG: "epilog",
	SHUTDOWN: "shutdown", CLEANUP_RESUBMIT: "cleanup_resubmit", UNKNOWN: "unknown", HOTPLUG: "hotplug",
	SHUTDOWN_POWEROFF: "shutdown_poweroff", BOOT_UNKNOWN: "boot_unknown", BOOT_POWEROFF: "boot_poweroff",
	BOOT_SUSPENDED: "boot_suspended", BOOT_STOPPED: "boot_stopped", CLEANUP_DELETE: "cleanup_delete",
	HOTPLUG_SNAPSHOT: "hotplug_snapshot", HOTPLUG_NIC: "hotplug_nic", HOTPLUG_SAVEAS: "hotplug_saveas",
	HOTPLUG_SAVEAS_POWEROFF: "hotplug_saveas_poweroff", HOTPLUG_SAVEAS_SUSPENDED: "hotplug_saveas_suspended",
	SHUTDOWN_UNDEPLOY: "shutdown_undeploy", EPILOG_UNDEPLOY: "epilog_undeploy", PROLOG_UNDEPLOY: "prolog_undeploy",
	BOOT_UNDEPLOY: "boot_undeploy", HOTPLUG_PROLOG_POWEROFF: "hotplug_prolog_poweroff",
	HOTPLUG_EPILOG_POWEROFF: "hotplug_epilog_poweroff", BOOT_MIGRATE: "boot_migrate", BOOT_FAILURE: "boot_failure",
	BOOT_MIGRATE_FAILURE: "boot_migrate_failure", PROLOG_MIGRATE_FAILURE: "prolog_migrate_failure",
	PROLOG_FAILURE: "prolog_failure", EPILOG_FAILURE: "epilog_failure", EPILOG_STOP_FAILURE: "epilog_stop_failure",
	EPILOG_UNDEPLOY_FAILURE: "epilog_undeploy_failure", PROLOG_MIGRATE_POWEROFF: "prolog_migrate_poweroff",
	PROLOG_MIGRATE_POWEROFF_FAILURE: "prolog_migrate_poweroff_failure", PROLOG_MIGRATE_SUSPEND: "prolog_migrate_suspend",
	PROLOG_MIGRATE_SUSPEND_FAILURE: "prolog_migrate_suspend_failure", BOOT_UNDEPLOY_FAILURE: "boot_undeploy_failure",
	BOOT_STOPPED_FAILURE: "boot_stopped_failure", PROLOG_RESUME_FAILURE: "prolog_resume_failure",
	PROLOG_UNDEPLOY_FAILURE: "prolog_undeploy_failure", DISK_SNAPSHOT_POWEROFF: "disk_snapshot_poweroff",
	DISK_SNAPSHOT_REVERT_POWEROFF: "disk_snapshot_revert_poweroff", DISK_SNAPSHOT_DELETE_POWEROFF: "disk_snapshot_delete_poweroff",
	DISK_SNAPSHOT_SUSPENDED: "disk_snapshot_suspended", DISK_SNAPSHOT_REVERT_SUSPENDED: "disk_snapshot_revert_suspended",
	DISK_SNAPSHOT_DELETE_SUSPENDED: "disk_snapshot_delete_suspended", DISK_SNAPSHOT: "disk_snapshot",
	DISK_SNAPSHOT_DELETE: "disk_snapshot_delete", PROLOG_MIGRATE_UNKNOWN: "prolog_migrate_unknown",
	PROLOG_MIGRATE_UNKNOWN_FAILURE: "prolog_migrate_unknown_failure"}

var VmStateString = map[VmState]string{INIT: "init", PENDING: "pending", HOLD: "hold", ACTIVE: "active", STOPPED: "stopped", SUSPENDED: "suspended", DONE: "done", UNKNOWNSTATE: "unknown", POWEROFF: "poweroff", UNDEPLOYED: "undeployed"}

const (

	//LcmState starts at 0
	LCM_INIT LcmState = iota
	PROLOG
	BOOT
	RUNNING
	MIGRATE
	SAVE_STOP
	SAVE_SUSPEND
	SAVE_MIGRATE
	PROLOG_MIGRATE
	PROLOG_RESUME
	EPILOG_STOP
	EPILOG
	SHUTDOWN
	CLEANUP_RESUBMIT
	UNKNOWN
	HOTPLUG
	SHUTDOWN_POWEROFF
	BOOT_UNKNOWN
	BOOT_POWEROFF
	BOOT_SUSPENDED
	BOOT_STOPPED
	CLEANUP_DELETE
	HOTPLUG_SNAPSHOT
	HOTPLUG_NIC
	HOTPLUG_SAVEAS
	HOTPLUG_SAVEAS_POWEROFF
	HOTPLUG_SAVEAS_SUSPENDED
	SHUTDOWN_UNDEPLOY
	EPILOG_UNDEPLOY
	PROLOG_UNDEPLOY
	BOOT_UNDEPLOY
	HOTPLUG_PROLOG_POWEROFF
	HOTPLUG_EPILOG_POWEROFF
	BOOT_MIGRATE
	BOOT_FAILURE
	BOOT_MIGRATE_FAILURE
	PROLOG_MIGRATE_FAILURE
	PROLOG_FAILURE
	EPILOG_FAILURE
	EPILOG_STOP_FAILURE
	EPILOG_UNDEPLOY_FAILURE
	PROLOG_MIGRATE_POWEROFF
	PROLOG_MIGRATE_POWEROFF_FAILURE
	PROLOG_MIGRATE_SUSPEND
	PROLOG_MIGRATE_SUSPEND_FAILURE
	BOOT_UNDEPLOY_FAILURE
	BOOT_STOPPED_FAILURE
	PROLOG_RESUME_FAILURE
	PROLOG_UNDEPLOY_FAILURE
	DISK_SNAPSHOT_POWEROFF
	DISK_SNAPSHOT_REVERT_POWEROFF
	DISK_SNAPSHOT_DELETE_POWEROFF
	DISK_SNAPSHOT_SUSPENDED
	DISK_SNAPSHOT_REVERT_SUSPENDED
	DISK_SNAPSHOT_DELETE_SUSPENDED
	DISK_SNAPSHOT
	DISK_SNAPSHOT_DELETE
	PROLOG_MIGRATE_UNKNOWN
	PROLOG_MIGRATE_UNKNOWN_FAILURE
)
const (
	//VmState starts at 0
	INIT VmState = iota
	PENDING
	HOLD
	ACTIVE
	STOPPED
	SUSPENDED
	DONE
	UNKNOWNSTATE
	POWEROFF
	UNDEPLOYED
)

type Query struct {
	VMName string
	VMId   int
	T      *api.Rpc
}

type UserVMs struct {
	XMLName xml.Name `xml:"VM_POOL"`
	UserVM  []*VM    `xml:"VM"`
}

type UserVM struct {
	Id   int    `xml:"ID"`
	Uid  int    `xml:"UID"`
	Name string `xml:"NAME"`
}

type Vnc struct {
	VmId string
	T    *api.Rpc
	VM   *VM `xml:"VM"`
}

type VM struct {
	XMLName        xml.Name        `xml:"VM"`
	Id             string          `xml:"ID"`
	Name           string          `xml:"NAME"`
	State          int             `xml:"STATE"`
	LcmState       int             `xml:"LCM_STATE"`
	VmTemplate     *VmTemplate     `xml:"TEMPLATE"`
	UserTemplate   UserTemplate    `xml:"USER_TEMPLATE"`
	HistoryRecords *HistoryRecords `xml:"HISTORY_RECORDS"`
	Snapshots      *Snapshots      `xml:"SNAPSHOTS"`
}

type VmTemplate struct {
	Graphics *Graphics `xml:"GRAPHICS"`
	Context  *Context  `xml:"CONTEXT"`
	Nics     []Nic     `xml:"NIC"`
}

type Nic struct {
	Network   string `xml:"NETWORK"`
	Id        string `xml:"NIC_ID"`
	IPaddress string `xml:"IP"`
	Mac       string `xml:"MAC"`
}

type Context struct {
	VMIP string `xml:"ETH0_IP"`
}

type HistoryRecords struct {
	History *History `xml:"HISTORY"`
}
type History struct {
	HostName string `xml:"HOSTNAME"`
}

type Graphics struct {
	Port string `xml:"PORT"`
}

type UserTemplate struct {
	Description        string `xml:"DESCRIPTION"`
	Error              string `xml:"ERROR"`
	Sched_Requirements string `xml:"SCHED_REQUIREMENTS"`
}

type Snapshots struct {
	DiskId   int        `xml:"DISK_ID"`
	Snapshot []Snapshot `xml:"SNAPSHOT"`
}

type Snapshot struct {
	Name string `xml:"NAME"`
	Id   int    `xml:"ID"`
	Size string `xml:"SIZE"`
}

func (v *Vnc) GetVm() (*VM, error) {
	intstr, _ := strconv.Atoi(v.VmId)
	args := []interface{}{v.T.Key, intstr}
	onevm, err := v.T.Call(api.VM_INFO, args)
	if err != nil {
		return nil, err
	}

	xmlVM := &VM{}
	if err = xml.Unmarshal([]byte(onevm), xmlVM); err != nil {
		return nil, err
	}
	return xmlVM, err
}

//have to release hold ips
func (v *Vnc) AttachNic(network, ip string) error {
	var forceIp string
	id, _ := strconv.Atoi(v.VmId)
	if len(ip) > 0 {
		forceIp = ", IP=\"" + ip + "\""
	}
	nic := "NIC = [ NETWORK=\"" + network + "\", NETWORK_UNAME=\"oneadmin\"" + forceIp + "]"
	args := []interface{}{v.T.Key, id, nic}
	_, err := v.T.Call(api.ONE_VM_ATTACHNIC, args)
	return err
}

func (v *Vnc) DetachNic(nic int) error {
	id, _ := strconv.Atoi(v.VmId)
	args := []interface{}{v.T.Key, id, nic}
	_, err := v.T.Call(api.ONE_VM_DETACHNIC, args)
	return err
}

func (u *VM) GetPort() string {
	return u.VmTemplate.Graphics.Port
}

func (u *VM) GetState() int {
	return u.State
}

func (u *VM) GetLcmState() int {
	return u.LcmState
}

func (u *VM) GetHostIp() string {
	return u.HistoryRecords.History.HostName
}

func (u *VM) GetVMIP() string {
	return u.VmTemplate.Context.VMIP
}

func (v *VM) StateString() string {
	return VmStateString[VmState(v.State)]
}

func (v *VM) Nics() []Nic {
	return v.VmTemplate.Nics
}

func (v *VM) LenSnapshots() int {
	if v.Snapshots != nil {
		return len(v.Snapshots.Snapshot)
	}
	return 0
}

func (v *VM) NetworkIdByIP(ip string) string {
	for _, n := range v.VmTemplate.Nics {
		if ip == n.IPaddress {
			return n.Id
		}
	}
	return ""
}

func (v *VM) LcmStateString() string {
	return LcmStateString[LcmState(v.LcmState)]
}

func (v *VM) IsFailure() bool {
	return strings.Contains(v.LcmStateString(), "failure")
}

func (v *VM) IsSnapshotReady() bool {
	return (v.State == int(ACTIVE) && v.LcmState == int(RUNNING)) || (v.State == int(POWEROFF) && v.LcmState == int(LCM_INIT))
}

// Given a name, this function will return the VM
func (v *Query) GetByName() ([]*VM, error) {
	args := []interface{}{v.T.Key, -2, -1, -1, -1}
	VMPool, err := v.T.Call(api.VMPOOL_INFO, args)
	if err != nil {
		return nil, err
	}

	xmlVM := UserVMs{}
	if err = xml.Unmarshal([]byte(VMPool), &xmlVM); err != nil {
		return nil, err
	}
	var matchedVM = make([]*VM, len(xmlVM.UserVM))

	for _, u := range xmlVM.UserVM {
		if u.Name == v.VMName {
			matchedVM[0] = u
		}
	}

	return matchedVM, nil

}

const (
	ONE_USER_LOGIN = "one.user.login"
	ENDPOINT       = "endpoint"
	USERID         = "userid"
	CURRENTUSER    = "username"
	PASSWORD       = "password"
)

type Rpc struct {
	Client xmlrpc.Client
	Key    string
}

func NewMyClientSSL(config map[string]string, transport http.RoundTripper) (*api.Rpc, error) {

	if !satisfied(config) {
		return nil, nil
	}

	client, err := xmlrpc.NewClient(config[ENDPOINT], transport)
	if err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("  > [one-go] connection response", "blue", "", "bold")+"%#v",client)
	log.Print(cmd.Colorfy("  > [one-go] connected", "blue", "", "bold")+" %s", config[ENDPOINT])

	return &api.Rpc{
		Client: *client,
		Key:    config[USERID] + ":" + config[PASSWORD]}, nil
}

type Sgu struct {
	endpoint     string
	username     string
	password     string
	usertoken    string
	sguconfig    map[string]string
	tlsConfig    *tls.Config
	transport    *http.Transport
	apirpcclient *api.Rpc
}

func (s *Sgu) Init() {
	conf := make(map[string]string)
	conf[api.ENDPOINT] = s.endpoint
	conf[api.USERID] = s.username
	conf[api.PASSWORD] = s.password
	s.sguconfig = conf
	log.Print(s.sguconfig)

	s.tlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	s.transport = &http.Transport{
		TLSClientConfig: s.tlsConfig,
	}
	log.Print(s.sguconfig, s.transport)
	var err error

	s.apirpcclient, err = NewMyClientSSL(s.sguconfig, s.transport)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Sgu) login(hvsconf HvsConf) (HvsConf, string, error) {
	var err error
	currentTime := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	currentTimeSec := time.Since(currentTime).Seconds()
	log.Print(s.apirpcclient.Key)
	loginargs := []interface{}{
		s.apirpcclient.Key,
		s.username,
		s.usertoken,
		3000,
		-1,
	}
	if currentTimeSec-hvsconf.lasttokentime < 2500 {
		s.apirpcclient.Key = hvsconf.token
		return hvsconf, s.apirpcclient.Key, nil
	} else {
		s.usertoken, err = s.apirpcclient.Call(ONE_USER_LOGIN, loginargs)
		if err != nil {
			log.Fatal("Error Sgu Login: ", err)
			return hvsconf, "", err
		}
		str := strings.Split(s.apirpcclient.Key, ":")
		s.apirpcclient.Key = str[0] + ":" + s.usertoken
		hvsconf.lasttokentime = currentTimeSec
		return hvsconf, s.apirpcclient.Key, nil
	}
}

func (s *Sgu) getVmList() map[string]*VM {
	vmPoolInfoArgs := []interface{}{
		s.apirpcclient.Key,
		-1,
		-1,
		-1,
		-1,
	}
	resp, err := s.apirpcclient.Call(api.VMPOOL_INFO, vmPoolInfoArgs)
	if err != nil {
		log.Fatal("Error vmpoolInfo: ", err)
	}
	log.Println(resp)
	var vms UserVMs
	err = xml.Unmarshal([]byte(resp), &vms)
	if err != nil {
		log.Println(err)
	}
	vmList := make(map[string]*VM)
	for _, vm := range vms.UserVM {
		vmList[vm.Id] = vm
	}
	return vmList
}

func (s *Sgu) on(cur VmInfo) {
	id, _ := strconv.Atoi(cur.CurrentMachine)

	ternOnArgs := []interface{}{
		s.apirpcclient.Key,
		"resume",
		id,
	}
	resp, err := s.apirpcclient.Call(api.ONE_VM_ACTION, ternOnArgs)
	if err != nil {
		log.Fatal("Error turn on: ", err)
	}
	log.Print(resp)
}

func (s *Sgu) off(cur VmInfo) {
	id, _ := strconv.Atoi(cur.CurrentMachine)
	ternOnArgs := []interface{}{
		s.apirpcclient.Key,
		"poweroff",
		id,
	}
	resp, err := s.apirpcclient.Call(api.ONE_VM_ACTION, ternOnArgs)
	if err != nil {
		log.Fatal("Error turn on: ", err)
	}
	log.Print(resp)
}

func (s *Sgu) reboot(cur VmInfo) {
	id, _ := strconv.Atoi(cur.CurrentMachine)
	ternOnArgs := []interface{}{
		s.apirpcclient.Key,
		"reboot",
		id,
	}
	resp, err := s.apirpcclient.Call(api.ONE_VM_ACTION, ternOnArgs)
	if err != nil {
		log.Fatal("Error turn on: ", err)
	}
	log.Print(resp)
}

type HvsConf struct {
	serverip      string
	serverport    string
	userName      string
	token         string
	lasttokentime float64
}

func (hvsconf *HvsConf) Save() {
	log.Print(hvsconf)
	b, err := json.Marshal(&hvsconf)
	if err != nil {
		log.Fatal("Error to save configuration file: ")
	}
	err = ioutil.WriteFile("hvsconf.json", b, 0644)
	if err != nil {
		log.Fatal("Error to create conf file: ", err)
	}
}

func (hvsconf *HvsConf) Load() bool {
	file, err := os.Open("hvsconf.json")
	if err != nil {
		log.Print("Error to load conf: ", err)
		return false
	} else {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&hvsconf)
		if err != nil {
			log.Panic(err)
		}
		log.Print(hvsconf.serverip)
		return true
	}
}

type VmInfo struct {
	CurrentMachine   string
	CurrMachineState string
}

func raw_connect(host string, port string) (bool, error) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return false, err
	}
	if conn != nil {
		defer conn.Close()
		return true, nil
	}
	return true, nil
}

func main() {
	var sguUser Sgu
	var vminfo VmInfo
	var vmList map[string]*VM
	var broker, sguport string
	var hvsconf HvsConf
	gtk.Init(nil)
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка err: ", err)
	}
	b.AddFromString(morda)
	///b.AddFromFile("Morda.glade")
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	obj, _ := b.GetObject("ErrorDialog")
	errorDialog := obj.(*gtk.Dialog)

	obj, _ = b.GetObject("errlbl")
	errLabel := obj.(*gtk.Label)

	obj, _ = b.GetObject("errOk")
	errOkButton := obj.(*gtk.Button)
	errOkButton.Connect("clicked", func() {
		errorDialog.Close()
	})

	obj, _ = b.GetObject("AdminCheckDialog")
	adminCheckDialog := obj.(*gtk.Dialog)

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	mcd, err := b.GetObject("MachineControlDialog")
	if err != nil {
		errorDialog.Show()
		errLabel.SetText("Ошибка загрузки окна: " + err.Error())
		log.Fatal("Ошибка загрузки окна мцд: ", err)
	}
	winMcd := mcd.(*gtk.Dialog)
	winMcd.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("vmList")
	vmListGUI := obj.(*gtk.ListStore)

	obj, _ = b.GetObject("reAuth")
	reAuthButton := obj.(*gtk.Button)
	reAuthButton.Connect("clicked", func() {
		adminCheckDialog.Show()
	})

	obj, _ = b.GetObject("turnOnOff")
	turnOnOff := obj.(*gtk.Button)
	turnOnOff.Connect("clicked", func() {
		if vminfo.CurrMachineState == "active" {
			sguUser.off(vminfo)
		} else {
			sguUser.on(vminfo)
		}
	})

	obj, _ = b.GetObject("reboot")
	reboot := obj.(*gtk.Button)
	reboot.Connect("clicked", func() {
		sguUser.reboot(vminfo)
	})

	obj, _ = b.GetObject("turnOffComp")
	turnOffComp := obj.(*gtk.Button)
	turnOffComp.Connect("clicked", func() {
		_ = exec.Command("poweroff")
	})

	obj, _ = b.GetObject("vmSelection")
	vmSelection := obj.(*gtk.TreeSelection)
	vmSelection.Connect("changed", func() {
		_, iter, _ := vmSelection.GetSelected()
		statevalue, err := vmListGUI.GetValue(iter, 1)
		if err != nil {
			return
		}
		value, err := vmListGUI.GetValue(iter, 3)
		if err != nil {
			return
		}
		strstate, _ := statevalue.GetString()
		str, err := value.GetString()
		if err != nil {
			return
		}
		vminfo.CurrentMachine = str
		vminfo.CurrMachineState = strstate
		if strstate == "active" {
			turnOnOff.SetLabel("Выключить")
		} else {
			turnOnOff.SetLabel("Включить")
		}
	})

	obj, _ = b.GetObject("connectTo")
	connectTo := obj.(*gtk.Button)
	connectTo.Connect("clicked", func() {
		spiceUri := "spice://" + vmList[vminfo.CurrentMachine].GetHostIp() + ":" + vmList[vminfo.CurrentMachine].GetPort()
		cmd := exec.Command("remote-viewer", "-f", spiceUri)
		cmd.Run()
	})
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	authDialog, err := b.GetObject("AuthorizationDialog")
	if err != nil {
		log.Fatal("Ошибка загрузки окна авторизации:", err)
	}
	winAuth := authDialog.(*gtk.Dialog)
	winAuth.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("LoginEntry")
	loginEntry := obj.(*gtk.Entry)
	loginEntry.Connect("changed", func() {
		sguUser.username, _ = loginEntry.GetText()
		hvsconf.userName = sguUser.username
	})

	obj, _ = b.GetObject("PasswordEntry")
	passwordEntry := obj.(*gtk.Entry)
	passwordEntry.Connect("changed", func() {
		sguUser.password, _ = passwordEntry.GetText()

	})

	obj, _ = b.GetObject("authDnext")
	authDNext := obj.(*gtk.Button)
	authDNext.Connect("clicked", func() {
		sguUser.Init()
		hvsconf, _, err := sguUser.login(hvsconf)
		if err != nil {
			errorDialog.Show()
			errLabel.SetText("Login error: " + err.Error())
			log.Fatal(err)
			return
		}
		var vmlTemp map[string]*VM
		go func() {
			for {
				vmList = sguUser.getVmList()
				if !reflect.DeepEqual(vmlTemp, vmList) {
					vmListGUI.Clear()
					vmlTemp = vmList
					for id, vm := range vmlTemp {
						iter := vmListGUI.Append()
						vmid := vm.Id
						err := vmListGUI.Set(iter, []int{0, 1, 3}, []interface{}{vm.Name, vm.StateString(), vmid})
						if err != nil {
							log.Fatal(err)
						}
						log.Print(id)
					}
					time.Sleep(5 * time.Second)
				} else {
					time.Sleep(5 * time.Second)
				}
			}
		}()
		hvsconf.Save()
		winMcd.Show()
		winAuth.Hide()
	})
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	startDialog, err := b.GetObject("StartDialog")
	if err != nil {
		log.Fatal("Ошибка err: ", err)
	}
	win := startDialog.(*gtk.Dialog)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("brokerEntry")
	brokerEntry := obj.(*gtk.Entry)
	brokerEntry.Connect("changed", func() {
		broker, _ = brokerEntry.GetText()
	})

	obj, _ = b.GetObject("portEntry")
	portEntry := obj.(*gtk.Entry)
	portEntry.Connect("changed", func() {
		sguport, _ = portEntry.GetText()
	})

	obj, _ = b.GetObject("adminPass")
	adminPassEntry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("repeatAdmPass")
	repeatAdmPassEntry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("saveCheckButton")
	saveCheckButton := obj.(*gtk.CheckButton)

	obj, _ = b.GetObject("startNext")
	startNext := obj.(*gtk.Button)
	startNext.Connect("clicked", func() {
		service := "MAX"
		us := "admin"
		ushost := "admhost"
		usport := "admport"

		pas, err := adminPassEntry.GetText()
		if err != nil {
			log.Fatal(err)
		}
		pas2, err2 := repeatAdmPassEntry.GetText()
		if err2 != nil {
			log.Fatal(err2)
		}
		conn, err := raw_connect(broker, sguport)
		if err != nil {
			errorDialog.Show()
			errLabel.SetText("Broker or is invalid: " + err.Error())
			return
		}
		fmt.Print(conn)
		if pas == pas2 && saveCheckButton.Activate() {
			err := keyring.Set(service, us, pas)
			if err != nil {
				errorDialog.Show()
				errLabel.SetText("Keyring: " + err.Error())
				log.Fatal(err)
			}
			err = keyring.Set(service, ushost, broker)
			if err != nil {
				errorDialog.Show()
				errLabel.SetText("Keyring: " + err.Error())
				log.Fatal(err)
			}
			err = keyring.Set(service, usport, sguport)
			if err != nil {
				log.Fatal(err)
			}
		}

		sguUser.endpoint = "https://" + broker + ":" + sguport + "/RPC2"
		hvsconf.serverip = broker
		hvsconf.serverport = sguport
		log.Print(sguUser.endpoint)

		winAuth.Show()
		loginEntry.DeleteText(0, -1)
		passwordEntry.DeleteText(0, -1)
		win.Hide()
	})
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	obj, _ = b.GetObject("rememberedPas")
	rememberedPas := obj.(*gtk.Entry)

	obj, _ = b.GetObject("acdNext")
	acdNext := obj.(*gtk.Button)
	acdNext.Connect("clicked", func() {
		pas1, err := rememberedPas.GetText()
		if err != nil {
			errorDialog.Show()
			errLabel.SetText("Password1: " + err.Error())
		}
		secret, err := keyring.Get("MAX", "admin")
		if err != nil {
			errorDialog.Show()
			errLabel.SetText("Secret: " + err.Error())
		}
		if pas1 == secret {

			win.Show()

			adminCheckDialog.Hide()
		} else {
			errorDialog.Show()
			errLabel.SetText("Пароль неправильный!")
			return
		}
	})

	tmp := hvsconf.Load()
	if tmp {
		brokerEntry.SetText(hvsconf.serverip)
		portEntry.SetText(hvsconf.serverport)
		loginEntry.SetText(hvsconf.userName)
	}
	secret, err := keyring.Get("MAX", "admin")
	log.Print("SECRET: ", secret)
	if err != nil {
		log.Print(err)
	}
	if len(secret) > 0 {
		broker, _ = keyring.Get("MAX", "admhost")
		sguport, _ = keyring.Get("MAX", "admport")
		sguUser.endpoint = "https://" + broker + ":" + sguport + "/RPC2"
		hvsconf.serverip = broker
		hvsconf.serverport = sguport
		log.Print(sguUser.endpoint)

		winAuth.Show()
		gtk.Main()
	} else {
		win.Show()
		gtk.Main()
	}
	// fmt.Print(secret)
	// win.Show()
	// gtk.Main()
}
func satisfied(c map[string]string) bool {
	return len(strings.TrimSpace(c[ENDPOINT])) > 0 &&
		len(strings.TrimSpace(c[USERID])) > 0 &&
		len(strings.TrimSpace(c[PASSWORD])) > 0
}
